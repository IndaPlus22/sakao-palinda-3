###Task 1.

1. What happens if you remove the go-command from the Seek call in the main function?

Ans: The program will run sequentially and make it so Anna will send a message to bob, Cody will send to Dave and so on as the array goes on.

2. What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?

Ans: wg will be an object of sync.WaitGroup and wg.done() will make the local copy to done but nothing else. Removing * means it makes a copy and not a reference and makes wg.done() not do anything.

3. What happens if you remove the buffer on the channel match?

Ans: Causes deadlock. There are no go routines that can recieve from the channel and therefore is blocked. The buffer alows it to be sent.

4. What happens if you remove the default-case from the case-statement in the main function?

Ans: No problem right now, but if there were even amount of people it would cause a deadlock because there would be messages left in the channel.

|Variant       | Runtime (ms) |
| ------------ | ------------:|
| singleworker |        28.59 |
| mapreduce    |        26.56 |