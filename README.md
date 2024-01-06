# Aule

All of my scripts for IRE Muds.

## AuleUI

Sets up a basic UI that supports IRE-based MUDs. If you wish to customize the layout, you can change the values in
`AuleSettings` if you write a handler for the `Aule.onInit` event:

```lua
function updateLayout()
    -- Update your settings in here.
end
registerNamedEventHandler("Aule", "updateLayout", "Aule.onInit", updateLayout)
```

**Setting the value to `0` will disable that side.**

Possible options are:

* `AuleSettings.left`: How wide the left panel should be
* `AuleSettings.right`: How wide the right panel should be
* `AuleSettings.top`: How tall the top panel should be
* `AuleSettings.bottom`: How tall the bottom panel should be

The top and bottom panels *always* stretch to fill the space between the left and right panels. If those panels are turned
off, they'll take the full width of the screen.