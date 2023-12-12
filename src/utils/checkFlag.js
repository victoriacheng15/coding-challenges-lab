export function checkFlag(args, obj) {
	const selectedFlag = args.find((flag) => obj.hasOwnProperty(flag));
	if (selectedFlag) {
		console.log(obj[selectedFlag]);
	}
}
