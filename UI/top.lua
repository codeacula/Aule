function Aule.setupTop()
    if AuleSettings.top == 0 then
        return
    end

    local top = Aule.container("top", {
        x = AuleSettings.left,
        y = 0,
        width = AuleSettings.clientWidth - AuleSettings.left - AuleSettings.right,
        height = AuleSettings.top,
    })
end