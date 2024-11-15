const config = {
  oidc: {
    authority: "...",
    client_id: "...",
    redirect_uri: `${window.location.origin}/redirect`,
  }
};

console.log(config);