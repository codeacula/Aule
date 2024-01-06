function Aule.setupRight()
    local right = Aule.container("right", {
        x = AuleSettings.clientWidth - AuleSettings.right,
        y = 0,
        width = AuleSettings.right,
        height = AuleSettings.clientHeight,
    });
end