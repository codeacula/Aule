function Aule.setupTop()
    local top = Aule.container("top", {
        x = AuleSettings.left,
        y = 0,
        width = AuleSettings.clientWidth - AuleSettings.left - AuleSettings.right,
        height = AuleSettings.top,
    })
end