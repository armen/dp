```
Algorithm 1.2: Asynchronous Job Handler

Implements:
   JobHandler, instance jh.

upon event < jh, Init > do
   buffer := ∅ ;

upon event < jh, Submit | job > do
   buffer := buffer ∪ { job } ;
   trigger < jh, Confirm | job > ;

upon buffer != ∅ do
   job := selectjob (buffer);
   process(job);
   buffer := buffer \ { job } ;
```
