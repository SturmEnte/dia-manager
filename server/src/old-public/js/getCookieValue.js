function getCookieValue(cookieName) {
	const cookieString = document.cookie;
	const regex = new RegExp("(^|;)\\s*" + cookieName + "\\s*=\\s*([^;]*)");
	const match = regex.exec(cookieString);
	return match ? decodeURIComponent(match[2]) : null;
}
