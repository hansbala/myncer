type AccessCodeResponseType = {
    access_token: string
    token_type: string
    expires_in: number
}

export const getSpotifyAccessToken = async (clientId: string, clientSecret: string): Promise<string> => {
    const url = 'https://accounts.spotify.com/api/token';
    const headers: HeadersInit = {
        'Content-Type': 'application/x-www-form-urlencoded',
    };
    const body: URLSearchParams = new URLSearchParams();
    body.append('grant_type', 'client_credentials');
    body.append('client_id', clientId);
    body.append('client_secret', clientSecret);

    const response: Response = await fetch(url, {
        method: 'POST',
        headers: headers,
        body: body,
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json() as AccessCodeResponseType
    return data.access_token
}