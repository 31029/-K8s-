package configs

import (
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8sapi/src/models"

	"k8sapi/src/services"
	"log"
)

type K8sConfig struct {
	DepHandler   *services.DepHandler     `inject:"-"`
	PodHandler   *services.PodHandler     `inject:"-"`
	NsHandler    *services.NsHandler      `inject:"-"`
	NodeHandler  *services.NodeMapHandler `inject:"-"`
	EventHandler *services.EventHandler   `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

//初始化客户端
func (*K8sConfig) InitClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	config.Insecure = true
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化 系统 配置
func (*K8sConfig) InitSysConfig() *models.SysConfig {
	b, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &models.SysConfig{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//初始化Informer
func (this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(this.PodHandler)

	nsInformer := fact.Core().V1().Namespaces() //监听namespace
	nsInformer.Informer().AddEventHandler(this.NsHandler)

	NodeInformer := fact.Core().V1().Nodes()
	NodeInformer.Informer().AddEventHandler(this.NodeHandler)

	eventInformer := fact.Core().V1().Events() //监听event
	eventInformer.Informer().AddEventHandler(this.EventHandler)
	fact.Start(wait.NeverStop)

	return fact
}
