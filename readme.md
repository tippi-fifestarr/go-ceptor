### Go Ceptor

Ceptor Club's mission is to improve the TTRPG experience with Art + Tech + Games. With this software, how can we make it easier, faster, and more fun to play together?

## Features

- [ ] Pregenerated characters from **Drive, Astrovan, Drive**
- [ ] Adventure hooks for easy GMing and unique experiences for each playthrough
- [ ] Information about each "starting location"
- [ ] A robust login and award tracking system
- [ ] User can set their availability and preferences
- [ ] Gamemaster can set up a game and invite players

## Lessons from Disco
- [ ] Best practices

Abstract a NewCharacter with defaults and for NewSpecific do a NewCharacter with custom variables
Safety first, check nils (or always intialize)
Don't access the struct directly, use getters 
Rounding up

**Abilities should be a Struct not a mapping**

- [ ] Quiz me

- [ ] Using GPT better

- [ ] Can this be in onebot?
- [ ] -- watch how onebot is structured
- [ ] -- try to seperate the core library features from the web features
-- make handleFunc an API endpoint for a react front-end