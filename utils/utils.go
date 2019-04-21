package utils

import (
	"encoding/json"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"strconv"

	"github.com/shinytang6/openpaas-security/client"
	deploy "github.com/shinytang6/openpaas-security/deployment"
	svc "github.com/shinytang6/openpaas-security/service"
	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var clientset *kubernetes.Clientset
var index int
//var globalpath string = ""

// 不能写在StartOneExperiment里面, 会报错, 所以改成导入包时就初始化client
func init() {
	clientset = client.InitClient()
	index = 0
}

func StartOneExperiment(config string) {
	index += 1
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
	serviceClient := clientset.CoreV1().Services("default")

	deployment := GenerateDeployment("../examples/dvwa.yaml")
	deploy.CreateDeployment(deploymentClient, deployment)
	service := GenerateService("../examples/dvwa_svc.yaml")
	svc.CreateService(serviceClient, service)

	vncdeployment := GenerateDeployment("../examples/vnc.yaml")
	vncdeployment = updateDeployment(vncdeployment, "vnc" + strconv.Itoa(index))
	deploy.CreateDeployment(deploymentClient, vncdeployment)

	vncservice := GenerateService("../examples/vnc_svc.yaml")
	vncservice = updateService(vncservice, "vnc" + strconv.Itoa(index))
	svc.CreateService(serviceClient, vncservice)
}

func GenerateDeployment(filename string) apps_v1beta1.Deployment {
	var (
		deployYaml []byte
		deployJson []byte
		deployment = apps_v1beta1.Deployment{}
		err error
	)

	if deployYaml, err = ioutil.ReadFile(filename); err != nil {
		panic(err.Error())
	}

	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		panic(err.Error())
	}
	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		panic(err.Error())
	}
	return deployment
}

func GenerateService(filename string) core_v1.Service {
	var (
		serviceYaml []byte
		serviceJson []byte
		service = core_v1.Service{}
		err error
	)

	if serviceYaml, err = ioutil.ReadFile(filename); err != nil {
		panic(err.Error())
	}

	if serviceJson, err = yaml.ToJSON(serviceYaml); err != nil {
		panic(err.Error())
	}

	if err = json.Unmarshal(serviceJson, &service); err != nil {
		panic(err.Error())
	}
	return service
}

func updateDeployment(deployment apps_v1beta1.Deployment, name string) apps_v1beta1.Deployment{
	deployment.ObjectMeta.Name = name
	deployment.Spec.Selector.MatchLabels = map[string]string{"app": name}
	deployment.Spec.Template.Labels = map[string]string{"app": name}
	deployment.Spec.Template.Spec.Containers[0].Name = name
	return  deployment
}

func updateService(svc core_v1.Service, name string) core_v1.Service {
	svc.ObjectMeta.Name = name
	svc.Spec.Selector = map[string]string{"app": name}
	return svc
}