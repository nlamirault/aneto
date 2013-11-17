help:
	@echo "release - package and upload a release"
	@echo "sdist   - package"
	@echo "doc     - generate HTML documentation"
	@echo "clean   - remove build artifacts"

release: docs
	rm -rf dist build
	python setup.py sdist upload

sdist: docs
	python setup.py sdist
	ls -l dist

clean:
	rm -rf dist build *.egg-info
	(cd doc && make clean)

.PHONY: doc
doc:
	(cd doc && make clean html)
