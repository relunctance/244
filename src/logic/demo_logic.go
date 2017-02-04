package logic

type DemoLogic struct {
	BaseLogic
}

/*
example:

	demoLogic := new(logic.DemoLogic)
	result, err := demoLogic.TestLogic("gaoqilin")
	if err != nil {
		fmt.Println(err, "###################\n")
	}
	fmt.Println(result)

*/

func (this *DemoLogic) TestLogic(a string) (string, error) {
	if a == "" {
		return "", this.throw_error(ERR_PARAM_FAILD, "params a is empty")
	}
	return a + ": success", nil
}
