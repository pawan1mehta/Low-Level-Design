# Error Group

An error group runs several goroutines and treats them as one unit: wait for all of them, and if any returns an error, that error is the result.
