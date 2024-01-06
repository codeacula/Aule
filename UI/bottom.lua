function Aule.setupBottom()
    if AuleSettings.bottom == 0 then
        return
    end

    local bottom = Aule.container("bottom", {
        x = AuleSettings.left,
        y = AuleSettings.clientHeight - AuleSettings.bottom,
        width = AuleSettings.clientWidth - AuleSettings.left - AuleSettings.right,
        height = AuleSettings.bottom,
    })
end