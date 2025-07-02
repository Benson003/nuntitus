<script lang="ts">

    import Container from "../components/base/Container.svelte";
    import Button from "../components/base/Button.svelte";
    import Inputs from "../components/base/inputs/TextInput.svelte";
    import Link from "../components/base/Links.svelte";

    import HXL from "../components/typography/HXL.svelte";
    import HL from "../components/typography/HL.svelte";
    import HM from "../components/typography/HM.svelte";
    import HSM from "../components/typography/HM.svelte";
    import PL from "../components/typography/PL.svelte";
    import PM from "../components/typography/PM.svelte";
    import PS from "../components/typography/PS.svelte";

    let demoInput = $state("");
</script>

    <Container class="space-y-8 bg-gray-700">

        <!-- Typography Demo -->
        <div class="space-y-2">
            <HXL class="text-blue-400 dark:text-blue-400">HXL: Hero Extra Large Heading</HXL>
            <HL class="text-red-400 dark:text-red-400">HL: Large Heading</HL>
            <HM class="text-green-400 dark:text-green-400">HM: Medium Heading</HM>
            <HSM>HSM: Small Heading</HSM>
            <PL>PL: Large Paragraph - Lorem ipsum dolor sit amet, consectetur adipiscing elit.</PL>
            <PM>PM: Medium Paragraph - This demonstrates paragraph styling in your design system.</PM>
            <PS>PS: Small Paragraph - Small, descriptive or hint text.</PS>
        </div>

        <!-- Inputs Demo -->
        <div class="space-y-4">
            <Inputs label="Demo Input" value={demoInput} type="text" class="dark:bg-gray-800"/>
            <Inputs label="Password" value={demoInput} type="password"class="dark:bg-gray-800" />
        </div>

        <!-- Buttons Demo -->
        <div class="space-y-2">
            <Button>Primary Button</Button>
            <Button class="bg-blue-600 hover:bg-blue-500">Blue Button</Button>
            <Button class="bg-green-600 hover:bg-green-500">Green Button</Button>
            <Button class="bg-red-600 hover:bg-red-500">Green Button</Button>
        </div>

        <!-- Link Demo -->
        <div class="text-center">
            <Link href="https://www.google.com"newTab >Visit Google</Link>
        </div>

    </Container>
