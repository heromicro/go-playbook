package main

import (
	"github.com/heromicro/go-playbook/builtin/opgather"
	"github.com/heromicro/go-playbook/play"
)

func main() {

	gf := opgather.AnsibleBuiltinGatherFacts{
		Name: "gather facts " + "localhost" + "[\"" + "127.0.0.1" + "\", ]",
	}

	pb := play.Playbook{
		Hosts:       []string{"127.0.0.1", ""},
		GatherFacts: enumtipe.CostomBoolFalse,
		Tasks:       []play.ITaskMaker{&gf},
	}

	ansible_result, pb_content, duration, err := pb.ExecPlaybook(context.TODO())
	// pb_content, err := pb.MakeAnsibleTask()
	fmt.Println(" === duration: ", duration)
	fmt.Println(" === pb_content: ", pb_content)
	if err != nil {
		panic(err)
	}

	r := helper.MarshalIndentToString(ansible_result, "", "  ")
	fmt.Println(" === result: ", r)

}
