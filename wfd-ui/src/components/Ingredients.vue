<template>
  <v-card>
    <v-card-title>
      Ingredients
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="6" v-for="(ingredient, index) in ingredientStrings" :key="index">
          <v-card>
            <v-card-text>
              {{ ingredient }}
            </v-card-text>
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

const ingredientStrings = computed(() => {
  return props.ingredients.map((ingredient) => {
    return `${ingredient.quantity} ${ingredient.unit} ${ingredient.name}`;
  });
});
</script>