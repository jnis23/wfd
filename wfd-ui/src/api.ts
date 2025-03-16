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


const api = new Api("http://localhost:8080");

export default api;