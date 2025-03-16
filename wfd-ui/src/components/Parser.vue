<template>
    <v-container>
        <v-row>
            <v-col cols="12">
                <v-text-field label="URL" v-model="url" />
            </v-col>
            <v-col cols="4">
                <v-btn @click="parse">Parse</v-btn>
            </v-col>
        </v-row>
        <v-row v-if="recipe">
            <v-col cols="12">
                <v-card>
                    <v-card-title>
                        {{ recipe.Name }}
                    </v-card-title>
                    <v-card-text>
                        {{ recipe.Description }}
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="12">
                <Ingredients :ingredients="recipe.Ingredients" />
            </v-col>
            <v-col cols="12">
                <Instructions :instructions="recipe.Instructions" />
            </v-col>
        </v-row>
    </v-container>
</template>

<script setup lang="ts">
import type { Recipe } from "@/types";
import api from "@/api";

const url = ref("");
const recipe = ref<Recipe>();

async function parse(){
    const res = await api.parse(url.value);
    recipe.value = res;
};
</script>