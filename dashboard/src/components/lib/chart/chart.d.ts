/**
 *
 * Chart components are based on Chart.js, an open source HTML5 based charting library.
 *
 * [Live Demo](https://www.primefaces.org/primereact/chart)
 *
 * @module chart
 *
 */
import * as React from 'react';

/**
 * Defines valid properties in Chart component.
 * @group Properties
 */
export interface ChartProps {
    /**
     * Unique identifier of the element.
     */
    id?: string | undefined;
    /**
     * Type of the chart.
     */
    type?: string | undefined;
    /**
     * Data to display.
     */
    data?: object | undefined;
    /**
     * Options to customize the chart.
     */
    options?: object | undefined;
    /**
     * Used to custom plugins of the chart.
     */
    plugins?: any[] | undefined;
    /**
     * Width of the chart in non-responsive mode.
     */
    width?: string | undefined;
    /**
     * Height of the chart in non-responsive mode.
     */
    height?: string | undefined;
    /**
     * Inline style of the element.
     */
    style?: React.CSSProperties | undefined;
    /**
     * Style class of the element.
     */
    className?: string | undefined;
    /**
     * Used to get the child elements of the component.
     * @readonly
     */
    children?: React.ReactNode | undefined;
}

/**
 * **PrimeReact - Chart**
 *
 * _Chart components are based on Chart.js, an open source HTML5 based charting library._
 *
 * [Live Demo](https://www.primefaces.org/primereact/chart/)
 * --- ---
 * ![PrimeReact](https://primefaces.org/cdn/primereact/images/logo-100.png)
 *
 * @group Component
 */
export declare class Chart extends React.Component<ChartProps, any> {
    /**
     * Used to get canvas element.
     * @return {HTMLCanvasElement} Canvas element
     */
    public getCanvas(): HTMLCanvasElement;
    /**
     * Used to get chart instance.
     * @return {*} Chart instance
     */
    public getChart(): any;
    /**
     * Used to get base64 image.
     * @return {*} base64 image
     */
    public getBase64Image(): any;
    /**
     * Used to generate legend.
     * @return {string} Generated legend
     */
    public generateLegend(): string;
    /**
     * Redraws the graph.
     */
    public refresh(): void;
    /**
     * Used to get container element.
     * @return {HTMLDivElement} Container element
     */
    public getElement(): HTMLDivElement;
}