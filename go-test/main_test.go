package main

import "testing"



//table driven test cases
type addTest struct{
	a1,a2,r int
}

var addTests=[]addTest{
	addTest{1,3,4},
	addTest{4,5,9},
	addTest{11,22,33},
	addTest{9,0,9},
}

func TestAdd(t *testing.T){
	for _,test:= range addTests{
		if output:=AddNum(test.a1, test.a2); output!=test.r{
			t.Errorf("output %v not equal to expected %v", output,test.r)
		}
	}
}

/*
Test case


func TestAdd(t *testing.T) {
	got := AddNum(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %v , wanted %v", got, want)
	}else{
		t.Log("correct answer")
	}
}

*/