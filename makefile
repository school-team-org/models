push:
	git add .
	git commit -m "$m"
	git push
	git tag v$m
	git push v$m