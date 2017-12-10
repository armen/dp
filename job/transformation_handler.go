package job

// Interface and properties of a job transformation and processing abstraction.
//
// Properties:
// 	TH1: Guarnteed response
// 		- Every submitted job is eventually confirmed or its transformation fails
// 	TH2: Soundness
// 		- A submitted job whose transformation fails is not processed
//
type TransformationHandler interface {
	Handler           // Inherits the Submit(...) and Confirm(...) from job.Handler
	Error(func(*Job)) // Indicates that the transformation of the given job failed
}
