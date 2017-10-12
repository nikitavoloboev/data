# Sticky Keys

Snippet below makes `a` key into a modifier key. This only works for old Karabiner on El Captain. I am still not certain how it will look for Sierra.

## sticky key example

\`\`\`XML
<item>
	<name>general: sticky p</name>
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
	  KeyCode::P,
	
	  <!--
	      Use notsave.launcher_mode_v2 in order to be higher priority.
	      Use ModifierFlag::MY_LAUNCHER_MODE for __DropAllKeys__.
	  -->
	  @begin
	  KeyCode::VK_CONFIG_SYNC_KEYDOWNUP_notsave_private_launcher_mode_v12, ModifierFlag::MY_LAUNCHER_MODE,
	  @end
	
	  @begin
	  KeyCode::P,
	  @end
	</autogen>
  </item>

  <item hidden="true">
	<identifier vk_config="true">notsave.private_launcher_mode_v12</identifier>
	
	<autogen>
	  __BlockUntilKeyUp__ KeyCode::P,
	</autogen>
	
	<autogen>
	    __KeyToKey__
	    KeyCode::A,
	    KeyCode::A, ModifierFlag::CONTROL_L
	</autogen>

  </item>
\`\`\`

This will make your P key into a modifier key where if you press it once, it will insert P. However if you press and _hold_ P key, it becomes your own personal modifier key. 

This code : 

\`\`\`XML
	<autogen>
	    __KeyToKey__
	    KeyCode::A,
	    KeyCode::A, ModifierFlag::CONTROL_L
	</autogen>
\`\`\`
Makes it so that pressing A key when P is pressed will make control + A action.

You can extend this to add more modifier key presses this way.

# Notes

- You can probably use [this](https://github.com/tekezo/Karabiner-Elements/issues/926) to achieve this behaviour on Sierra but I have not tested it
e