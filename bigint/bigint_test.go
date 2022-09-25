package bigint

import (
	"testing"

)

func TestNewInt(t *testing.T){
    tests := []struct{
        num      string
        expected string
        err       error
      
    }{
        {num:"123", expected: "123"},
        {num:"0", expected: "0"}, 
        {num:"-628", expected: "-628"},
        {num:"00087", expected: "87"},
        {num:"-00087", expected: "-87"},
        {num:"+87", expected: "87"},
     //   {num:"+8-7", err:ErrorLength},
        
     
        
   
    }
    for _,  tt := range tests{
        res, _ := NewInt(tt.num)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input value: %v",res.Value,tt.expected,tt.num)
    
        }
    }
     t.Log("Finished")
}

func TestSet(t *testing.T){
    tests := []struct{
        num Bigint
        expected error
    }{
        {Bigint{Value: "343"},nil},
         {Bigint{Value: "-343"}, nil},
        {Bigint{Value: "++10"}, ErrorLength},
       {Bigint{Value: "4375464865754753735343"}, ErrorLength},
    }
    for _, tt := range tests{
        res := tt.num.Set(tt.num.Value)
        if res != tt.expected{
            t.Errorf("got %v but expected %v, input value: %v", res, tt.expected, tt.num.Value)
        }
    }
    t.Log("Finished")
}
    

func TestAdd(t *testing.T){

    tests := []struct{
        num1, num2 Bigint
        expected string
    }{
     {num1:Bigint{Value:"123"}, num2:Bigint{"634"}, expected: "757"},
     {num1:Bigint{Value:"36347"}, num2:Bigint{"46346"}, expected: "82693"},
     {num1:Bigint{Value:"34647"}, num2:Bigint{"5353"}, expected: "40000"},
     {num1:Bigint{Value:"0"}, num2:Bigint{"0"}, expected: "0"},
     {num1:Bigint{Value:"-235325"}, num2:Bigint{"-345"}, expected: "-235670"},
     {num1:Bigint{Value:"-123"}, num2:Bigint{"634"}, expected: "511"},
     {num1:Bigint{Value:"123"}, num2:Bigint{"-634"}, expected: "-511"},
    }
    
    for _,  tt := range tests{
        res := Add(tt.num1, tt.num2)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input values: %v, %v",res.Value,tt.expected,tt.num1.Value,tt.num2.Value)
    
        }
    }
     t.Log("Finished")
}

func TestSub(t *testing.T){

    tests := []struct{
        num1, num2 Bigint
        expected string
    }{
     {num1:Bigint{Value:"123"}, num2:Bigint{"634"}, expected: "-511"},
     {num1:Bigint{Value:"36347"}, num2:Bigint{"46346"}, expected: "-9999"},
     {num1:Bigint{Value:"34647"}, num2:Bigint{"5353"}, expected: "29294"},
     {num1:Bigint{Value:"0"}, num2:Bigint{"0"}, expected: "0"},
     {num1:Bigint{Value:"-235325"}, num2:Bigint{"-345"}, expected: "-234980"},
     {num1:Bigint{Value:"-123"}, num2:Bigint{"634"}, expected: "-757"},
     {num1:Bigint{Value:"123"}, num2:Bigint{"-634"}, expected: "757"},
    }
    
    for _,  tt := range tests{
        res := Sub(tt.num1, tt.num2)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input values: %v, %v",res.Value,tt.expected,tt.num1.Value,tt.num2.Value)
        }
    }
     t.Log("Finished")
}

func TestMultiply(t *testing.T){

    tests := []struct{
        num1, num2 Bigint
        expected string
    }{
     {num1:Bigint{Value:"123"}, num2:Bigint{"634"}, expected: "77982"},
     {num1:Bigint{Value:"36347"}, num2:Bigint{"46346"}, expected: "1684538062"},
     {num1:Bigint{Value:"0"}, num2:Bigint{"0"}, expected: "0"},
     {num1:Bigint{Value:"-235325"}, num2:Bigint{"-345"}, expected: "81187125"},
     {num1:Bigint{Value:"-123"}, num2:Bigint{"634"}, expected: "-77982"},
     {num1:Bigint{Value:"123"}, num2:Bigint{"-634"}, expected: "-77982"},
    }
    
    for _,  tt := range tests{
        res := Multiply(tt.num1, tt.num2)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input values: %v, %v",res.Value,tt.expected,tt.num1.Value,tt.num2.Value)
        }
    }
     t.Log("Finished")
}
func TestMod(t *testing.T){

    tests := []struct{
        num1, num2 Bigint
        expected string
    }{
     {num1:Bigint{Value:"123"}, num2:Bigint{"634"}, expected: "123"},
     {num1:Bigint{Value:"34647"}, num2:Bigint{"5353"}, expected: "2529"},
     {num1:Bigint{Value:"0"}, num2:Bigint{"0"}, expected: "Division by is undifined "},
     {num1:Bigint{Value:"-235325"}, num2:Bigint{"-345"}, expected: "-35"},

    }
    
    for _,  tt := range tests{
        res := Mod(tt.num1, tt.num2)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input values: %v, %v",res.Value,tt.expected,tt.num1.Value,tt.num2.Value)
        }
    }
     t.Log("Finished")
}
func TestDivede(t *testing.T){

    tests := []struct{
        num1, num2 Bigint
        expected string
    }{
     {num1:Bigint{Value:"123"}, num2:Bigint{"634"}, expected: "0"},
     {num1:Bigint{Value:"36347"}, num2:Bigint{"46346"}, expected: "0"},
     {num1:Bigint{Value:"34647"}, num2:Bigint{"5353"}, expected: "6"},
     {num1:Bigint{Value:"0"}, num2:Bigint{"0"}, expected: "Division by is undifined "},
     {num1:Bigint{Value:"-235325"}, num2:Bigint{"-345"}, expected: "682"},
     {num1:Bigint{Value:"-123"}, num2:Bigint{"634"}, expected: "0"},
     {num1:Bigint{Value:"123"}, num2:Bigint{"-634"}, expected: "0"},
    }
    
    for _,  tt := range tests{
        res := Divide(tt.num1, tt.num2)
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input values: %v, %v",res.Value,tt.expected,tt.num1.Value,tt.num2.Value)
        }
    }
     t.Log("Finished")
}


func TestAbs(t *testing.T){
    tests := []struct{
        num Bigint
        expected string
    }{
        {num:Bigint{Value: "343"}, expected: "343"},
        {num:Bigint{Value: "-343"}, expected: "343"},
    }
    for _, tt := range tests{
        res := tt.num.Abs()
        if res.Value != tt.expected{
            t.Errorf("got %v but expected %v, input value: %v", res.Value, tt.expected, tt.num.Value)
        }
    }
    t.Log("Finished")
}