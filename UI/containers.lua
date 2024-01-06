function Aule.container(name, settings, parent)
    settings.name = name

    local container = Geyser.Container:new(settings, parent)

    AuleWindows[name] = container

    return container
end

function Aule.map(name, settings, parent)
    settings.name = name

    local map = Geyser.Mapper:new(settings, parent)

    AuleWindows[name] = map

    return map
end