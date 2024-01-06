function Aule.setupRight()
    if AuleSettings.right == 0 then
        return
    end

    local right = Aule.container("right", {
        x = AuleSettings.clientWidth - AuleSettings.right,
        y = 0,
        width = AuleSettings.right,
        height = AuleSettings.clientHeight,
    })
end