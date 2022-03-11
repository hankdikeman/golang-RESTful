package main

// import string conversion and Beego
import (
    "strconv"
    "github.com/astaxie/beego"
)

// driver function for instantiating routes
func main() {
    // generic path for math operation
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})

    // start server
    beego.Run()
}

// declare struct from beego request controller
type mainController struct {
    beego.Controller
}

// GET route for handling math ops
func (c *mainController) Get() {
    
    // retrieve inputs from route parameters above
    operation := c.Ctx.Input.Param(":operation")
    num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
    num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

    // set values for input into template
    c.Data["operation"] = operation
    c.Data["num1"] = num1
    c.Data["num2"] = num2
    c.TplName = "result.html"

    // provide result based on operation parameter
    switch operation {
        case "sum":
            c.Data["result"] = add(num1, num2)
        case "product":
            c.Data["result"] = multiply(num1, num2)
        default:
            c.TplName = "invalid-route.html"
    }
}

func add(n1, n2 int) int {
    return n1 + n2
}

func multiply(n1, n2 int) int {
    return n1 * n2
}
