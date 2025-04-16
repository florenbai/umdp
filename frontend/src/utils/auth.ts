const TOKEN_KEY = 'umdp_session';

// 获取Cookie函数
const getCookie = (name: string) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) {
    return parts.pop()?.split(';').shift();
  }
  return null;
};

// 判断Cookie是否过期的函数
const isCookieExpired = (cookieName: string) => {
  const cookie = getCookie(cookieName);
  if (!cookie) {
    return true; // 如果Cookie不存在，则认为已过期
  }
  const expires = new Date(cookie.split('=')[1]); // 解析出过期时间
  return expires < new Date(); // 如果过期时间小于当前时间，则认为已过期
};

const deleteCookie = (cookieName: string) => {
  document.cookie = `${cookieName}=; expires=Thu, 01 Jan 1970 00:00:01 GMT;`;
};

const isLogin = () => {
  return !isCookieExpired(TOKEN_KEY);
};

const getToken = () => {
  return getCookie(TOKEN_KEY);
};

const setToken = (token: string) => {
  localStorage.setItem(TOKEN_KEY, token);
  const expirationDate = new Date();
  expirationDate.setDate(expirationDate.getDate() + 30);
  document.cookie = `${TOKEN_KEY}=${token}; expires=${expirationDate.toUTCString()}; path=/`;
};

const clearToken = () => {
  deleteCookie(TOKEN_KEY);
};

export { isLogin, getToken, setToken, clearToken };
