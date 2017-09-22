
# Sticky Keys

```XML
<item>
		<name>general: sticky a - control key</name>
		<identifier>private.launcher_mode_open_apps_v12</identifier>

    <!-- condition: turn on launcher mode only when the trigger key is pressed without other keys. -->
    <pressingphysicalkeys_lessthan>2</pressingphysicalkeys_lessthan>
    <modifier_not>
      ModifierFlag::COMMAND_L,
      ModifierFlag::COMMAND_R,
      ModifierFlag::CONTROL_L,
      ModifierFlag::CONTROL_R,
      ModifierFlag::FN,
      ModifierFlag::OPTION_L,
      ModifierFlag::OPTION_R,
      ModifierFlag::SHIFT_L,
      ModifierFlag::SHIFT_R,
    </modifier_not>

    <autogen>
      __KeyOverlaidModifier__
      KeyCode::A,

      <!--
          Use notsave.launcher_mode_v2 in order to be higher priority.
          Use ModifierFlag::MY_LAUNCHER_MODE for __DropAllKeys__.
      -->
      @begin
      KeyCode::VK_CONFIG_SYNC_KEYDOWNUP_notsave_private_launcher_mode_v93, ModifierFlag::MY_LAUNCHER_MODE,
      @end

      @begin
      KeyCode::A,
      @end
    </autogen>
  </item>

  <item hidden="true">
    <identifier vk_config="true">notsave.private_launcher_mode_v12</identifier>

    <autogen>
      __BlockUntilKeyUp__ KeyCode::A,
    </autogen>

    <autogen>
        __KeyToKey__
        KeyCode::F,
        KeyCode::F, ModifierFlag::CONTROL_L, 
    </autogen>

  </item>
```

This makes `a` key into a modifier key. 