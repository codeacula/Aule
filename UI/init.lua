function Aule.init()
    raiseEvent("Aule.onInit")
    AuleSettings.clientWidth, AuleSettings.clientHeight = getMainWindowSize()

    for key, window in pairs(AuleWindows) do
        window:hide()
        AuleWindows[key] = nil
    end

    -- Make space for the UI
    setBorderLeft(AuleSettings.left)
    setBorderTop(AuleSettings.top)
    setBorderRight(AuleSettings.right)
    setBorderBottom(AuleSettings.bottom)

    Aule.setupRight()
    Aule.setupLeft()
    Aule.setupTop()
    Aule.setupBottom()
end

Aule.init()