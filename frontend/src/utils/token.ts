const getTokenHeader = () => {
  const token = sessionStorage.getItem("token");
  return `Bearer ${token}`;
};

export default getTokenHeader;
