<template>
  <b-card class="header" header-bg-variant="primary">
        <template #header>
            <h4 class="mb-0">Author</h4>
        </template>
        <b-list-group>
            <b-list-group-item id="user" class="d-flex align-items-center">
                <b-avatar style="float:left" variant="secondary" class="mr-4"></b-avatar>
                <span style="float:right" class="mr-auto">J. Circlehead</span>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <span style="float:left" class="mr-auto">Last login on: </span>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <span style="float:left" class="mr-auto">Location: </span>
            </b-list-group-item>
        </b-list-group>
    </b-card>
</template>

<script>
export default {
    props: ["userid"]
}
</script>

<style>
#user:hover {
    background-color: var(--secondary);
    cursor:pointer;
    color:white;
}
</style>