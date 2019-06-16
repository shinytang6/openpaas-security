package utils

import (
	"encoding/json"
	"fmt"
	"github.com/shinytang6/openpaas-security/client"
	deploy "github.com/shinytang6/openpaas-security/deployment"
	svc "github.com/shinytang6/openpaas-security/service"
	"io/ioutil"
	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"time"
)

var clientset *kubernetes.Clientset
var index int
//var globalpath string = ""

// 不能写在StartOneExperiment里面, 会报错, 所以改成导入包时就初始化client
// 若要持久化最好将index存进数据库，这里偷懒了
func init() {
	clientset = client.InitClient()
	index = 0
}

func StartOneExperiment(config string, name string) {
	index += 1
	// 启动实验是根据yaml配置文件来创建的, 一般一个实验需要启动一个deployment资源以及service资源, 目前项目中所有试验相关配置文件都在backend/examples下
	// 未来应该将配置文件抽离出来, 单独放在一个地方维护
	deployment_config_file := "../examples/" + config + ".yaml"
	service_config_file := "../examples/" + config + "_svc" + ".yaml"
	fmt.Println(deployment_config_file, service_config_file)
	// 初始化客户端
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
	serviceClient := clientset.CoreV1().Services("default")

	//deployment := GenerateDeployment("../examples/dvwa.yaml")
	deployment := GenerateDeployment(deployment_config_file)
	deployment = updateDeployment(deployment, name)
	deploy.CreateDeployment(deploymentClient, deployment)
	//service := GenerateService("../examples/dvwa_svc.yaml")
	service := GenerateService(service_config_file)
	service = updateService(service, name)
	svc.CreateService(serviceClient, service)

	// 创建vnc depolyment和service资源
	// 每一个实验都需要vnc远程桌面
	vncdeployment := GenerateDeployment("../examples/vnc.yaml")
	vncdeployment = updateDeployment(vncdeployment, "vnc-" + name)
	deploy.CreateDeployment(deploymentClient, vncdeployment)

	vncservice := GenerateService("../examples/vnc_svc.yaml")
	vncservice = updateService(vncservice, "vnc-" + name)
	svc.CreateService(serviceClient, vncservice)
}

func DeleteOneExperiment(name string) {
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
	serviceClient := clientset.CoreV1().Services("default")

	deploy.DeleteDeployment(deploymentClient, name)
	svc.DeleteService(serviceClient, name)
	deploy.DeleteDeployment(deploymentClient, "vnc-" + name)
	svc.DeleteService(serviceClient, "vnc-" + name)
}

func RestartOneExperiment(name string) {
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
	serviceClient := clientset.CoreV1().Services("default")

	deploy.DeleteDeployment(deploymentClient, "vnc-" + name)
	svc.DeleteService(serviceClient, "vnc-" + name)

	time.Sleep(5 * time.Second)

	vncdeployment := GenerateDeployment("../examples/vnc.yaml")
	vncdeployment = updateDeployment(vncdeployment, "vnc-" + name)
	deploy.CreateDeployment(deploymentClient, vncdeployment)

	vncservice := GenerateService("../examples/vnc_svc.yaml")
	vncservice = updateService(vncservice, "vnc-" + name)
	svc.CreateService(serviceClient, vncservice)
}

func GetOneExperiment(name string) *core_v1.Service {
	serviceClient := clientset.CoreV1().Services("default")

	svc, _ := svc.GetService(serviceClient, "vnc-" + name)
	return svc
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