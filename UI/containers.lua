function Aule.container(name, settings, parent)
    settings.name = name

    local container = Geyser.Container:new({
        settings
    }, parent)

    AuleWindows[name] = container

    return container
end