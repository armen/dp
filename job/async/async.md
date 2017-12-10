**Algorithm 1.2**: Asynchronous Job Handler

**Implements**:

	JobHandler, **instance** jh.

**upon event** ^C jh, Init ^D do

	buffer := ∅ ;

**upon event** ^C jh, Submit | job ^D do

	buffer := buffer ∪ { job } ;

	**trigger** ^C jh, Confirm | job ^D ;

**upon** buffer ^G = ∅ do

	job := selectjob (buffer);

	process(job);

	buffer := buffer \ { job } ;
