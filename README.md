# bank
A common dice game we will be using for a bot fighting tournament.

## Rules
- Everybody starts with 0 points. 
- There are 20 rounds. 
- Every round begins with a pot of 0. 
- Every time you roll the dice, add whatever you get to the pot. 
- Rolling a 1 ends the round and empties the pot.
- Rolling a 2 doubles whatever is currently in the pot.  
- If your first roll is a 1 or a 2, add what you got to the pot instead of following the previous two rules.

Every time the dice is rolled you will be giving a chance to "bank" (see title of game). Banking means you add the
current score of the pot to your personal score, but you can only bank once per round. The object of the game is to end
with more points than your opponents.

## Other Rules

No using reflection to access anything private from the `bank` package.

## Implementing

Implement the `bank.PlayerStrategy` interface found in `player.go`. See comments on that interface for details.

Create your bot in the example package, or any other package of your choice that isn't the bank package.

## Logistics

We will be having a training on this repo in March. When we actually do the tournament, we will do so by running 
1,000,000 games and whoever has the highest win percentage will move on.

When we do this we will do it in a separate repo so don't make any pull requests to this repo unless they are for fixes
to the logic of the game. PLEASE DON'T COMMIT YOUR BOTS HERE because then everybody will see them and see how your bot
plays and steal your logic and you'll lose.
