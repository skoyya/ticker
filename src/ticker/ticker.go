package ticker

import ("fmt"
	"time"
)

const  HOURS_TO_RUN = 3 
type UpdateTicker struct {
	tickerName  string
	tickerValue string
}

type Ticker struct {
	secondTicker string
	minuteTicker string
	hourTicker   string
	isRunning bool
	updateChannel  chan UpdateTicker
}

func  NewTicker() *Ticker {
	return &Ticker{"tick", "tock", "bong", false, make(chan UpdateTicker)}
}

func (t *Ticker) updateTickerValue(upd UpdateTicker) {
	switch upd.tickerName {
		case "secondTicker":  t.secondTicker = upd.tickerValue
		case "minuteTicker":  t.minuteTicker = upd.tickerValue
		case "hourTicker":  t.hourTicker = upd.tickerValue
	default:
		fmt.Println("Invalid Key " + upd.tickerName)
	}
}

func (t *Ticker) UpdateTicker(tickerName string, tickerValue string) {
	u := UpdateTicker{tickerName, tickerValue}
	t.updateChannel <- u
}

func (t *Ticker) IsRunning() bool {
	return t.isRunning;
}

func (t *Ticker) RunTicker() {
	//Trigger every second
    ticker := time.NewTicker(1 * time.Second)
	t.isRunning = true
    go func() {
		counter := 0
		hours := 0
        for {
           select {
				case <- ticker.C:
					counter++
					if counter == 360 {
						fmt.Println(counter, t.hourTicker)
						counter = 0
						hours++
						if hours == HOURS_TO_RUN {
							close(t.updateChannel)
							t.isRunning = false
							ticker.Stop()
							return
						}
					}else if counter%60 == 0 {
						fmt.Println(counter, t.minuteTicker)
					}else {
						fmt.Println(counter, t.secondTicker)
					}
				case upd := <-t.updateChannel:
					t.updateTickerValue(upd)
            }
        }
	}()
}
