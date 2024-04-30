package dataService

import (
	cm "IOT/modules/dataService/mqtt"
	"IOT/modules/dataService/tcp"
	tphttp "IOT/others/http"
	"IOT/services"
	uuid "IOT/utils"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()
	reg_mqtt_root()
	listenMQTT()
	// listenTCP()
}

func loadConfig() {
	log.Println("read config")
	var err error
	envConfigFile := flag.String("config", "./modules/dataService/config.yml", "path of configuration file")
	flag.Parse()
	viper.SetConfigFile(*envConfigFile)
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("FAILURE", err)
		return
	}
}

func listenMQTT() {
	var TSKVS services.TSKVService
	mqttHost := os.Getenv("TP_MQTT_HOST")
	if mqttHost == "" {
		mqttHost = viper.GetString("mqtt.broker")
	}
	broker := mqttHost
	uuid := uuid.GetUuid()
	clientid := viper.GetString(uuid)
	user := viper.GetString("mqtt.user")
	pass := viper.GetString("mqtt.pass")
	p, _ := ants.NewPool(500)
	pOther, _ := ants.NewPool(500)
	cm.Listen(broker, user, pass, clientid, func(c mqtt.Client, m mqtt.Message) {
		_ = p.Submit(func() {
			TSKVS.MsgProc(m.Payload(), m.Topic())
		})
	}, func(c mqtt.Client, m mqtt.Message) {
		_ = pOther.Submit(func() {
			TSKVS.MsgProcOther(m.Payload(), m.Topic())
		})
	}, func(c mqtt.Client, m mqtt.Message) {
		_ = p.Submit(func() {
			TSKVS.GatewayMsgProc(m.Payload(), m.Topic())
		})
	})

}

// 废弃
func ListenTCP() {
	tcpPort := viper.GetString("tcp.port")
	log.Printf("config of tcp port -- %s", tcpPort)
	tcp.Listen(tcpPort)
}
func reg_mqtt_root() {
	MqttHttpHost := os.Getenv("MQTT_HTTP_HOST")
	if MqttHttpHost == "" {
		MqttHttpHost = viper.GetString("api.http_host")
	}
	resps, errs := tphttp.Post("http://"+MqttHttpHost+"/v1/accounts/root", "{\"password\":\""+viper.GetString("mqtt.pass")+"\"}")
	if errs != nil {
		log.Println("Response1:", errs.Error())
	} else {
		defer resps.Body.Close()
		if resps.StatusCode == 200 {
			body, errs := ioutil.ReadAll(resps.Body)
			if errs != nil {
				log.Println("Response2:", errs.Error())
			} else {
				log.Println("Response3: ", string(body))
			}
		} else {
			log.Println("Get failed with error:" + resps.Status)
		}
	}
}
