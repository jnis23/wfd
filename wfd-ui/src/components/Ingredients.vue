<template>
  <v-card>
    <v-card-title>
      Ingredients
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="6" v-for="ingredient in ingredients" :key="ingredient.name">
          <v-card>
            <v-card-text>{{ ingredient.name }}</v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import type { Ingredient } from "@/types";

const props = defineProps<{
  ingredients: Ingredient[];
}>();
</script>