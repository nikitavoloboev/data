# Sticky Keys
### [sticky key example](https://gist.github.com/b591b290c6a55ac47b19158c721415a4)
Snippet above makes `a` key into a modifier key. This only works for old Karabiner on El Captain. I am still not certain how it will look like in Sierra.

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
- You can probably use [this](https://github.com/tekezo/Karabiner-Elements/issues/926) to achieve this behaviour on Sierra but I have not tested it.