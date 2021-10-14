<template>
    <b-card class="mb-3" header-bg-variant="primary" 
                    style="min-width: 20rem;"

    >
        <template #header>
            <h4 class="mb-0">Description</h4>
        </template>
        <b-list-group>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.price}}€</pre></span>
                    <span style="float: left">
                        <pre>Price: </pre></span>
                </p>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.color}}</pre></span>
                    <span style="float: left">
                        <pre>Color: </pre></span>
                </p>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.condition}}</pre></span>
                    <span style="float: left">
                        <pre>Condition: </pre></span>
                </p>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.size}}</pre></span>
                    <span style="float: left">
                        <pre>Size: </pre></span>
                </p>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.location}}</pre></span>
                    <span style="float: left">
                        <pre>Location: </pre></span>
                </p>
            </b-list-group-item>
            <b-list-group-item class="list-group-item border-0">
                <p>
                    <span style="float: right" class="description-vars">
                        <pre>{{product.createdAt}}</pre></span>
                    <span style="float: left">
                        <pre>Post created:  </pre></span>
                </p>
            </b-list-group-item>
        </b-list-group>
    </b-card>
</template>

<script>
    export default {
        props: ["product"]
    }
</script>

<style>
pre {
    vertical-align: middle;
}
</style>