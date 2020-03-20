package ticker

import "testing"


func TestUpdateTickerValue(t *testing.T) {
	tk := NewTicker()
	want := "TEST_secondTicker"
	tk.updateTickerValue(UpdateTicker{"secondTicker", want})

	if tk.secondTicker != want  {
		t.Errorf("updateTickerValue failed, expected[%s] but got [%s]",want,tk.secondTicker)
	}
}

func TestIsRunning(t *testing.T) {
	want := false
	tk := NewTicker()
	if tk.isRunning != want {
		t.Errorf("isRunning is failed, expected[%t] but got [%t]",want,tk.isRunning)
	}
}

func TestUpdateTicker(t *testing.T) {
	tk := NewTicker()
	want := false
	got := tk.UpdateTicker("secondTicker", "secondTicker")

	if got != want  {
		t.Errorf("UpdateTicker is failed, expected[%t] but got [%t]", want, got)
	}
}

func TestRunTicker(t *testing.T) {
	tk := NewTicker()
	want := true
	tk.RunTicker()
	if tk.isRunning != want {
		t.Errorf("RunTicker is failed, expected[%t] but got [%t]", want, tk.isRunning)
	}
}
