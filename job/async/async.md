**Algorithm 1.1**: Synchronous Job Handler
**Implements**:
	JobHandler, **instance** jh.

**upon event** < jh, Submit | job > **do**
	process(job);
	**trigger** < jh, Confirm | job >;
