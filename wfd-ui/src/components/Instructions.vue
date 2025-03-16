<template>
    <v-card>
        <v-card-title>
            Instructions
        </v-card-title>
        <v-card-text>
            <v-row>
                <v-col cols="12" v-for="(instruction, index) in instructions" :key="index">
                    <v-card>
                        <v-card-text>{{ instruction.text }}</v-card-text>
                    </v-card>
                </v-col>
            </v-row>

        </v-card-text>
    </v-card>
</template>

<script setup lang="ts">
import type { Instruction } from "@/types";

const props = defineProps<{
    instructions: Instruction[];
}>();
</script>