function Aule.setupLeft()
    if AuleSettings.left == 0 then
        return
    end
    
    local left = Aule.container("left", {
        x = 0,
        y = 0,
        width = AuleSettings.left,
        height = AuleSettings.clientHeight,
    })
end