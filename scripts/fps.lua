speed = 50
log("set initial speed to " .. speed)

log("start video stream")
startvideo()

repeat
    lastkey = getlastkeypress()

    if lastkey == "a" then
        left(speed)
    elseif lastkey == "d" then
        right(speed)
    elseif lastkey == "w" then
        forward(speed)
    elseif lastkey == "s" then
        backward(speed)
    elseif lastkey == "arrowup" then
        up(speed)
    elseif lastkey == "arrowdown" then
        down(speed)
    elseif lastkey == "arrowleft" then
        rotateleft(speed)
    elseif lastkey == "arrowright" then
        rotateright(speed)
    elseif lastkey == "q" then
        leftflip()
    elseif lastkey == "e" then
        rightflip()
    elseif lastkey == "r" then
        backflip()
    elseif lastkey == "f" then
        frontflip()
    elseif lastkey == "b" then
        bounce()
    elseif lastkey == "o" then
        speed = speed - 5
        if speed < 0 then
            speed = 0
        end
        log("reduce speed to " .. speed)
    elseif lastkey == "p" then
        speed = speed + 5
        if speed > 100 then
            speed = 100
        end
        log("increase speed to " .. speed)
    elseif lastkey == "z" then
        log("initiating landing")
        land()
        log("landed")
    elseif lastkey == "x" then
        log("initiating take off")
        takeoff()
        log("hovering")
    elseif lastkey == nil or lastkey == "" then
        -- do nothing
    else
        stop()
    end
until lastkey == "escape"

log("execution terminated")