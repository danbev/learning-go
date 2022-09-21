
out/%: src/%.go | out
	go build -o $@ $<

out/%_test: src/%_test.go | out
	go test -o $@ $<

out:
	@mkdir out

.PHONY: clean
clean:
	@${RM} -rf out
