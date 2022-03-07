package main

import "testing"

type TestValidityTest struct {
	arg      string
	expected bool
}

var TestValidityTests = []TestValidityTest{
	TestValidityTest{"1-TQyPu-8827630-XaU-3031460318-RRLVewyoB-9-FUEbuBlPW", true},
	TestValidityTest{"7-WYpQ-4062402-HxWhhC-89-tMkS-8217607-AhoiD-29969122-INfg-40981-woljMtyWA-7941138844-WEwwh-02064-ahBaj-4001207-PCxb-5095-cFX", true},
	TestValidityTest{"2wUbrX-97-dKYrBcL-40612-G", false},
}

func TestTestValidity(t *testing.T) {

	for _, test := range TestValidityTests {
		if output := TestValidity(test.arg); output != test.expected {
			t.Errorf("Output is not equal to expected")
		}
	}
}

type AverageNumberTest struct {
	arg      string
	expected float32
}

var AverageNumberTests = []AverageNumberTest{
	AverageNumberTest{"1-TQyPu-8-XaU-30-RRLVewyoB-9-FUEbuBlPW", 12},
	AverageNumberTest{"7-WYpQ-40-HxWhhC-8-tMkS-82-AhoiD-29-INfg-4-woljMtyWA-7-WEwwh-64-ahBaj-4-PCxb-50-cFX", 29},
}

func TestAverageNumber(t *testing.T) {

	for _, test := range AverageNumberTests {
		if output := AverageNumber(test.arg); output != test.expected {
			t.Errorf("Output is not equal to expected")
		}
	}
}

type WholeStoryTest struct {
	arg, expected string
}

var WholeStoryTests = []WholeStoryTest{
	WholeStoryTest{"1-TQyPu-8-XaU-30-RRLVewyoB-9-FUEbuBlPW", "TQyPu XaU RRLVewyoB FUEbuBlPW"},
	WholeStoryTest{"7-WYpQ-40-HxWhhC-8-tMkS-82-AhoiD-29-INfg-4-woljMtyWA-7-WEwwh-64-ahBaj-4-PCxb-50-cFX", "WYpQ HxWhhC tMkS AhoiD INfg woljMtyWA WEwwh ahBaj PCxb cFX"},
}

func TestWholeStory(t *testing.T) {

	for _, test := range WholeStoryTests {
		if output := WholeStory(test.arg); output != test.expected {
			t.Errorf("Output is not equal to expected")
		}
	}
}

type StoryStatTest struct {
	arg, expected1, expected2, expected3 string
	expected4                            []string
}

var StoryStatTests = []StoryStatTest{
	StoryStatTest{"1-hello-2-world", "hello", "hello", "5", []string{"hello", "world"}},
 

func TestStoryStat(t *testing.T) {

	for _, test := range StoryStatTests {
		if output1, output2, output3, output4 := StoryStat(test.arg); output1 != test.expected1 || output2 != test.expected2 || output3 != test.expected3 || len(output4) == 0 {
			t.Errorf("Output is not equal to expected")
		}
	}
}
