package lmjaeger

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/logicmonitor/k8s-argus/pkg/config"
	"github.com/logicmonitor/k8s-argus/pkg/lmctx"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	//"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"gopkg.in/yaml.v2"
)

var (
	lmconf *config.Config
)

func Initialise(lmconfig *config.Config) {
	conf := readConf()
	lmconf = lmconfig
	cl, _ := conf.InitGlobalTracer("argus")
	shutdownHook(cl)
}

func shutdownHook(closer io.Closer) {
	go func(cl io.Closer) {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Warn("Closing tracer on shutdown signal")
		cl.Close()
	}(closer)
}

func readConf() *jaegercfg.Configuration {
	configBytes, err := ioutil.ReadFile("/etc/argus/jaeger-config.yaml")
	if err != nil {
		log.Fatalf("Failed to read jaeger config file: /etc/argus/jaeger-config.yaml")
	}
	log.Debugf("jaeger config raw: %s", configBytes)
	conf := &jaegercfg.Configuration{}
	err = yaml.Unmarshal(configBytes, conf)
	if err != nil {
		log.Fatalf("Couldn't parse jaeger-config.yaml: %s", err.Error())
	}
	log.Infof("jaeger config read: %+v", conf)
	conf.Reporter.AttemptReconnectInterval = time.Duration(10 * time.Second)
	return conf
}

type LMSpan interface {
	opentracing.Span
	Info(alternatingKeyValues ...interface{})
	Error(alternatingKeyValues ...interface{})
	K8sObject(key string, object interface{})
}

type LMSpanObject struct {
	opentracing.Span
}

func (span *LMSpanObject) Info(alternatingKeyValues ...interface{}) {
	alternatingKeyValues = append(alternatingKeyValues, "level")
	alternatingKeyValues = append(alternatingKeyValues, "info")
	span.LogKV(alternatingKeyValues...)
}
func (span *LMSpanObject) K8sObject(key string, object interface{}) {
	objJson, err := json.Marshal(object)
	if err != nil {
		span.LogEventWithPayload(key, object)
	} else {
		span.LogEventWithPayload(key, string(objJson))
	}
}

func (span *LMSpanObject) Error(alternatingKeyValues ...interface{}) {
	alternatingKeyValues = append(alternatingKeyValues, "level")
	alternatingKeyValues = append(alternatingKeyValues, "error")
	span.LogKV(alternatingKeyValues...)
	span.SetTag("error", true)
}

type SpanReseter func()

// StartSpan creates new span and injects it in lmcontext object
// consumer must call reset - preferrably defer reset function to reset previous span in lm context
func StartSpan(lctx *lmctx.LMContext, operationName string, options ...opentracing.StartSpanOption) (LMSpan, SpanReseter) {
	parentSpan := Span(lctx)
	span := opentracing.StartSpan(operationName, options...)
	lmSpanObj := &LMSpanObject{
		Span: span.(opentracing.Span),
	}
	logger := lctx.Extract("logger")
	if logger != nil {
		lmSpanObj.SetBaggageItem("debug_id", fmt.Sprintf("%v", logger.(*log.Entry).Data["debug_id"]))
		lmSpanObj.SetTag("debug_id", lmSpanObj.BaggageItem("debug_id"))
	}
	lmSpanObj.SetTag("argus.cluster.name", lmconf.ClusterName)
	lctx.Set("span", lmSpanObj)
	if parentSpan != nil {
		return lmSpanObj, func() {
			lctx.Set("span", parentSpan)
		}
	}
	return lmSpanObj, func() {}
}

// Span returns span object from lmcontext object
func Span(lctx *lmctx.LMContext) LMSpan {
	span := lctx.Extract("span")
	if span == nil {
		return nil
	}
	return span.(LMSpan)
}
