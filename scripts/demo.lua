-- Simple demo script with some flips and turns

log("initiating take off")
takeoff()
log("hovering")

sleep(1000)

log("fly in a circle")
forward(20)
rotateleft(70)
sleep(6500)
stop()

sleep(1000)

log("front flip")
frontflip()
log("back flip")
backflip()

log("fly in a circle")
backward(20)
rotateright(70)
sleep(6500)
stop()

sleep(1000)

log("left flip")
leftflip()
log("right flip")
rightflip()

log("initiating landing")
land()
log("landed")