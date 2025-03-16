import type { Recipe } from "@/types";

class Api {
    private baseUrl: string;

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    async parse(url: string): Promise<Recipe> {
        const response = await fetch(`${this.baseUrl}/parse?url=${url}`);
        return response.json();
    }
}


const host = window.location.hostname;

const api = new Api(`http://${host}:8080`);

export default api;