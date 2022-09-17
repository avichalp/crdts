#### CRDTs
Conflict-free Replicated Data Types



### Run tests 
```go
go test ./...
```

### run benchmarks for vector clocks
```go
go test ./... -bench=.
```

#### References

###### CRDTs
- https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type
- http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-order-theory/
- http://jtfmumm.com/blog/2015/11/24/crdt-primer-2-convergent-crdts/
- https://josephg.com/blog/crdts-go-brrr/
- https://martin.kleppmann.com/papers/crdt-isabelle-oopsla17.pdf
- https://hal.inria.fr/inria-00555588/document
- https://developers.soundcloud.com/blog/roshi-a-crdt-system-for-timestamped-events


##### Vector Clocks (Logical time keeping)
- https://sookocheff.com/post/time/vector-clocks/
- http://archagon.net/blog/2018/03/24/data-laced-with-history/
- https://queue.acm.org/detail.cfm?id=2917756
- keeping time in real systems by kavya joshi https://www.youtube.com/watch?v=BRvj8PykSc4
- Hybrid clocks and population protocols: https://www.youtube.com/watch?v=YqNGbvFHoKM
- https://www.youtube.com/watch?v=x7drE24geUw
- https://www.youtube.com/watch?v=B5NULPSiOGw
- https://www.youtube.com/watch?v=4i7KrG1Afbk&list=UU_QIfHvN9auy2CoOdSfMWDw&index=596
- https://www.cockroachlabs.com/blog/living-without-atomic-clocks/


