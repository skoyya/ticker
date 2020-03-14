--- To change the ticker value using the curl command
curl -X PUT 'http://localhost:8088?tickerName=secondTicker&tickerValue=SSSSS' 
curl -X PUT 'http://localhost:8088?tickerName=minuteTicker&tickerValue=TTTTT' 
curl -X PUT 'http://localhost:8088?tickerName=hourTicker&tickerValue=HHHHH' 


Few things knowingly handled differently as I have concentrated on the acheiving the functionality

1) If the progra needs to be extended for 2Hrs ticker and 3Hrs ticker ..etc then I would use the single channel one reader just reads the message and prints. Reader do no know the timings.  Writer will have logic to handle what to write to the channel depending on the time tick. 

3) Ticker stops are harded coded 3 Hors, instead the same curl command can be used to control the ticker run time too
4) http server, continuosly runs. Keeping in the mind that we can restart the ticker using curl command again.
5) As the ticker values updation is happing in single go routin, concurrency is handled bydefault. No need to explicitely use mutex locks
6)unit test not wrote, production issue cameup and i got to concentrate on that. I can go over on the calls. 


