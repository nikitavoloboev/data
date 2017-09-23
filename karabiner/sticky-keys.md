
# Sticky Keys

Snippet below makes `a` key into a modifier key. This only works for old Karabiner on El Captain. I am still not certain how it will look for Sierra.

```XML
<item>
<name>general: sticky a - control key</name>
<identifier>private.launcher\_mode\_open\_apps\_v12</identifier>

<!-- condition: turn on launcher mode only when the trigger key is pressed without other keys. -->
\<pressingphysicalkeys\_lessthan\>2\</pressingphysicalkeys\_lessthan\>
\<modifier\_not\>
  ModifierFlag::COMMAND\_L,
  ModifierFlag::COMMAND\_R,
  ModifierFlag::CONTROL\_L,
  ModifierFlag::CONTROL\_R,
  ModifierFlag::FN,
  ModifierFlag::OPTION\_L,
  ModifierFlag::OPTION\_R,
  ModifierFlag::SHIFT\_L,
  ModifierFlag::SHIFT\_R,
\</modifier\_not\>

<autogen>
  __KeyOverlaidModifier__
  KeyCode::A,

  <!--

  Use notsave.launcher_mode_v2 in order to be higher priority.

  Use ModifierFlag::MY_LAUNCHER_MODE for __DropAllKeys__.

  -->
  @begin
  KeyCode::VK\_CONFIG\_SYNC\_KEYDOWNUP\_notsave\_private\_launcher\_mode\_v93, ModifierFlag::MY\_LAUNCHER\_MODE,
  @end

  @begin
  KeyCode::A,
  @end
</autogen>
  </item>

  <item hidden="true">
\<identifier vk\_config="true"\>notsave.private\_launcher\_mode\_v12</identifier>

<autogen>
  __BlockUntilKeyUp__ KeyCode::A,
</autogen>

<autogen>
__KeyToKey__
KeyCode::F,
KeyCode::F, ModifierFlag::CONTROL\_L, 
</autogen>

  </item>
```