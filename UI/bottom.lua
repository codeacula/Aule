function Aule.setupBottom()
    local bottom = Aule.container("bottom", {
        x = AuleSettings.left,
        y = AuleSettings.clientHeight - AuleSettings.bottom,
        width = AuleSettings.clientWidth - AuleSettings.left - AuleSettings.right,
        height = AuleSettings.bottom,
    })
end