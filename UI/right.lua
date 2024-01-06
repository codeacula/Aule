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

    local chat = Aule.container("chat", {
        x = 0, y = 0, width = "100%", height = "50%"
    }, right)

    local map = Aule.container("map", {
        x = 0, y = "50%", width = "100%", height = "50%"
    }, right)

    Aule.setupChat(chat)
    Aule.setupMap(map)
end